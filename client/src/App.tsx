import { useState, useEffect } from 'react';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { toast } from 'sonner';
import {
  Plus,
  Trash2,
  Edit3,
  Check,
  X,
  CheckCircle2,
  Circle,
  ListTodo,
  Target,
  Clock,
  Calendar,
  TrendingUp
} from 'lucide-react';

const API_BASE = 'http://localhost:8000/api';

interface Todo {
  _id: string;
  body: string;
  completed: boolean;
}

interface ApiResponse {
  success: boolean;
  message: string;
  data: Todo[] | Todo | null;
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState('');
  const [editingId, setEditingId] = useState<string | null>(null);
  const [editText, setEditText] = useState('');
  const [loading, setLoading] = useState(false);
  const [submitting, setSubmitting] = useState(false);

  // Fetch todos
  const fetchTodos = async () => {
    setLoading(true);
    try {
      const response = await fetch(`${API_BASE}/todos`);
      const data: ApiResponse = await response.json();
      if (data.success && Array.isArray(data.data)) {
        setTodos(data.data);
      }
    } catch (error) {
      toast.error('Failed to fetch todos');
    } finally {
      setLoading(false);
    }
  };

  // Create todo
  const createTodo = async () => {
    if (!newTodo.trim()) return;

    setSubmitting(true);
    try {
      const response = await fetch(`${API_BASE}/todos`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ body: newTodo.trim() }),
      });

      const data: ApiResponse = await response.json();
      if (data.success) {
        setNewTodo('');
        fetchTodos();
        toast.success('Task created successfully', {
          description: `"${newTodo.trim()}" added to your workflow`,
        });
      }
    } catch (error) {
      toast.error('Failed to create task', {
        description: 'Please check your connection and try again',
      });
    } finally {
      setSubmitting(false);
    }
  };

  // Update todo
  const updateTodo = async (id: string) => {
    if (!editText.trim()) return;

    try {
      const response = await fetch(`${API_BASE}/todos/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ body: editText.trim() }),
      });

      if (response.ok) {
        setEditingId(null);
        setEditText('');
        fetchTodos();
        toast.success('Task updated successfully', {
          description: 'Changes saved to your workflow',
        });
      }
    } catch (error) {
      toast.error('Failed to update task', {
        description: 'Please try again',
      });
    }
  };

  // Complete todo
  const completeTodo = async (id: string) => {
    try {
      const response = await fetch(`${API_BASE}/todos/${id}`, {
        method: 'PATCH',
      });

      if (response.ok) {
        fetchTodos();
        toast.success('Task completed!', {
          description: 'Great work on finishing this task',
        });
      }
    } catch (error) {
      toast.error('Failed to complete task', {
        description: 'Please try again',
      });
    }
  };

  // Delete todo
  const deleteTodo = async (id: string) => {
    try {
      const response = await fetch(`${API_BASE}/todos/${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        fetchTodos();
        toast.success('Task removed', {
          description: 'Task deleted from your workflow',
        });
      }
    } catch (error) {
      toast.error('Failed to delete task', {
        description: 'Please try again',
      });
    }
  };

  const startEdit = (todo: Todo) => {
    setEditingId(todo._id);
    setEditText(todo.body);
  };

  const cancelEdit = () => {
    setEditingId(null);
    setEditText('');
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const completedCount = todos.filter(todo => todo.completed).length;
  const totalCount = todos.length;
  const pendingCount = totalCount - completedCount;
  const completionRate = totalCount > 0 ? Math.round((completedCount / totalCount) * 100) : 0;

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-white to-slate-100">
      {/* Animated Background Elements */}
      <div className="fixed inset-0 overflow-hidden pointer-events-none">
        <div className="absolute -top-40 -right-40 w-80 h-80 bg-gradient-to-br from-blue-100 to-indigo-200 rounded-full opacity-20 animate-pulse"></div>
        <div className="absolute -bottom-40 -left-40 w-80 h-80 bg-gradient-to-tr from-slate-100 to-gray-200 rounded-full opacity-20 animate-pulse" style={{ animationDelay: '2s' }}></div>
      </div>

      <div className="relative z-10 max-w-5xl mx-auto p-6">
        {/* Header Section */}
        <div className="text-center mb-12 animate-fade-in">
          <div className="flex items-center justify-center gap-4 mb-8">
            <div className="relative">
              <div className="w-16 h-16 bg-gradient-to-br from-slate-800 to-slate-900 rounded-2xl flex items-center justify-center shadow-xl transform hover:scale-105 transition-transform duration-300">
                <ListTodo className="w-8 h-8 text-white" />
              </div>
              <div className="absolute -inset-2 bg-gradient-to-r from-blue-600 to-indigo-600 rounded-2xl blur opacity-20 animate-pulse"></div>
            </div>
            <div>
              <h1 className="text-6xl font-extralight text-slate-800 tracking-tight mb-2">
                TaskFlow
              </h1>
              <div className="h-1 w-24 bg-gradient-to-r from-slate-400 to-slate-600 rounded-full mx-auto"></div>
            </div>
          </div>

          <p className="text-xl text-slate-600 font-light mb-8 max-w-2xl mx-auto leading-relaxed">
            Streamline your productivity with intelligent task management
          </p>

          {/* Enhanced Stats Dashboard */}
          {totalCount > 0 && (
            <div className="grid grid-cols-1 md:grid-cols-4 gap-4 max-w-4xl mx-auto mb-8">
              <div className="bg-white/80 backdrop-blur-sm rounded-2xl p-6 shadow-lg border border-slate-200/50 hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1">
                <div className="flex items-center justify-between mb-2">
                  <Target className="w-6 h-6 text-slate-600" />
                  <span className="text-3xl font-light text-slate-800">{totalCount}</span>
                </div>
                <p className="text-slate-600 text-sm font-medium">Total Tasks</p>
              </div>

              <div className="bg-white/80 backdrop-blur-sm rounded-2xl p-6 shadow-lg border border-slate-200/50 hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1">
                <div className="flex items-center justify-between mb-2">
                  <CheckCircle2 className="w-6 h-6 text-green-600" />
                  <span className="text-3xl font-light text-green-700">{completedCount}</span>
                </div>
                <p className="text-slate-600 text-sm font-medium">Completed</p>
              </div>

              <div className="bg-white/80 backdrop-blur-sm rounded-2xl p-6 shadow-lg border border-slate-200/50 hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1">
                <div className="flex items-center justify-between mb-2">
                  <Clock className="w-6 h-6 text-amber-600" />
                  <span className="text-3xl font-light text-amber-700">{pendingCount}</span>
                </div>
                <p className="text-slate-600 text-sm font-medium">In Progress</p>
              </div>

              <div className="bg-white/80 backdrop-blur-sm rounded-2xl p-6 shadow-lg border border-slate-200/50 hover:shadow-xl transition-all duration-300 transform hover:-translate-y-1">
                <div className="flex items-center justify-between mb-2">
                  <TrendingUp className="w-6 h-6 text-blue-600" />
                  <span className="text-3xl font-light text-blue-700">{completionRate}%</span>
                </div>
                <p className="text-slate-600 text-sm font-medium">Progress</p>
              </div>
            </div>
          )}
        </div>

        {/* Task Input Section */}
        <div className="mb-8 animate-slide-up">
          <Card className="bg-white/90 backdrop-blur-sm border-0 shadow-2xl rounded-3xl overflow-hidden">
            <CardContent className="p-8">
              <div className="flex gap-4">
                <div className="relative flex-1">
                  <Input
                    placeholder="What would you like to accomplish today?"
                    value={newTodo}
                    onChange={(e) => setNewTodo(e.target.value)}
                    onKeyPress={(e) => e.key === 'Enter' && createTodo()}
                    className="h-14 text-lg bg-slate-50/50 border-slate-200 rounded-2xl px-6 focus:ring-2 focus:ring-slate-400 focus:border-transparent transition-all duration-300 placeholder:text-slate-400"
                    disabled={submitting}
                  />
                  <div className="absolute inset-0 bg-gradient-to-r from-blue-500/5 to-indigo-500/5 rounded-2xl pointer-events-none opacity-0 transition-opacity duration-300 peer-focus:opacity-100"></div>
                </div>
                <Button
                  onClick={createTodo}
                  disabled={!newTodo.trim() || submitting}
                  className="h-14 px-8 w-14 bg-gradient-to-r from-slate-800 to-slate-900 hover:from-slate-700 hover:to-slate-800 text-white rounded-2xl font-medium shadow-lg hover:shadow-xl transition-all duration-300 transform hover:scale-105 disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none"
                >
                  {submitting ? (
                    <div className="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin" />
                  ) : (
                    <Plus className="w-5 h-5" />
                  )}
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Tasks List Section */}
        <Card className="bg-white/90 backdrop-blur-sm border-0 shadow-2xl rounded-3xl overflow-hidden animate-slide-up">
          <CardHeader className="border-b border-slate-100 bg-gradient-to-r from-slate-50 to-white p-8">
            <div className="flex items-center justify-between">
              <div>
                <CardTitle className="text-2xl font-light text-slate-800 flex items-center gap-3">
                  <Calendar className="w-6 h-6" />
                  Your Tasks
                </CardTitle>
                <CardDescription className="text-slate-600 mt-2">
                  {totalCount === 0
                    ? "Ready to start your productive day?"
                    : `Managing ${totalCount} task${totalCount !== 1 ? 's' : ''} in your workflow`
                  }
                </CardDescription>
              </div>
              {totalCount > 0 && (
                <div className="hidden md:flex items-center gap-2 px-4 py-2 bg-slate-100 rounded-full">
                  <div className="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                  <span className="text-sm text-slate-600 font-medium">
                    {completionRate}% Complete
                  </span>
                </div>
              )}
            </div>
          </CardHeader>

          <CardContent className="p-0">
            {loading ? (
              <div className="flex items-center justify-center py-20">
                <div className="relative">
                  <div className="w-12 h-12 border-4 border-slate-200 rounded-full"></div>
                  <div className="absolute inset-0 w-12 h-12 border-4 border-slate-600 border-t-transparent rounded-full animate-spin"></div>
                </div>
              </div>
            ) : todos.length === 0 ? (
              <div className="text-center py-20 px-8">
                <div className="w-24 h-24 bg-gradient-to-br from-slate-100 to-slate-200 rounded-full flex items-center justify-center mx-auto mb-6 shadow-inner">
                  <ListTodo className="w-12 h-12 text-slate-400" />
                </div>
                <h3 className="text-2xl font-light text-slate-700 mb-3">No tasks yet</h3>
                <p className="text-slate-500 max-w-md mx-auto leading-relaxed">
                  Start building your productive workflow by adding your first task above.
                  Every great achievement begins with a single step.
                </p>
              </div>
            ) : (
              <div className="divide-y divide-slate-100">
                {todos.map((todo, index) => (
                  <div
                    key={`${todo._id}-${index}`}
                    className={`p-6 transition-all duration-300 hover:bg-slate-50/50 group ${todo.completed ? 'opacity-60' : ''
                      }`}
                    style={{
                      animationDelay: `${index * 100}ms`,
                      animation: 'fadeInUp 0.5s ease-out forwards'
                    }}
                  >
                    <div className="flex items-center gap-6">
                      {/* Completion Button */}
                      <Button
                        variant="ghost"
                        size="sm"
                        onClick={() => !todo.completed && completeTodo(todo._id)}
                        disabled={todo.completed}
                        className="p-0 h-auto hover:bg-transparent group/button"
                      >
                        <div className="relative">
                          {todo.completed ? (
                            <div className="w-8 h-8 bg-green-500 rounded-full flex items-center justify-center shadow-lg">
                              <Check className="w-5 h-5 text-white" />
                            </div>
                          ) : (
                            <div className="w-8 h-8 border-2 border-slate-300 rounded-full flex items-center justify-center group-hover/button:border-green-500 group-hover/button:bg-green-50 transition-all duration-200">
                              <div className="w-0 h-0 bg-green-500 rounded-full group-hover/button:w-3 group-hover/button:h-3 transition-all duration-200"></div>
                            </div>
                          )}
                        </div>
                      </Button>

                      {/* Task Content */}
                      <div className="flex-1 min-w-0">
                        {editingId === todo._id ? (
                          <div className="flex gap-3 animate-fade-in">
                            <Input
                              value={editText}
                              onChange={(e) => setEditText(e.target.value)}
                              onKeyPress={(e) => {
                                if (e.key === 'Enter') updateTodo(todo._id);
                                if (e.key === 'Escape') cancelEdit();
                              }}
                              className="bg-white border-slate-300 rounded-xl focus:ring-2 focus:ring-slate-400 focus:border-transparent"
                              autoFocus
                            />
                            <Button
                              size="sm"
                              onClick={() => updateTodo(todo._id)}
                              className="bg-green-600 hover:bg-green-700 text-white rounded-xl px-4 shadow-lg hover:shadow-xl transition-all duration-200"
                            >
                              <Check className="w-4 h-4" />
                            </Button>
                            <Button
                              size="sm"
                              variant="outline"
                              onClick={cancelEdit}
                              className="border-slate-300 text-slate-600 hover:bg-slate-50 rounded-xl px-4"
                            >
                              <X className="w-4 h-4" />
                            </Button>
                          </div>
                        ) : (
                          <p
                            className={`text-lg leading-relaxed transition-all duration-300 ${todo.completed
                                ? 'line-through text-slate-400'
                                : 'text-slate-700 group-hover:text-slate-900'
                              }`}
                          >
                            {todo.body}
                          </p>
                        )}
                      </div>

                      {/* Action Buttons */}
                      <div className="flex gap-2 opacity-0 group-hover:opacity-100 transition-all duration-200">
                        {!todo.completed && editingId !== todo._id && (
                          <Button
                            size="sm"
                            variant="ghost"
                            onClick={() => startEdit(todo)}
                            className="text-slate-500 hover:text-blue-600 hover:bg-blue-50 rounded-xl p-2 transition-all duration-200"
                          >
                            <Edit3 className="w-4 h-4" />
                          </Button>
                        )}
                        <Button
                          size="sm"
                          variant="ghost"
                          onClick={() => deleteTodo(todo._id)}
                          className="text-slate-500 hover:text-red-600 hover:bg-red-50 rounded-xl p-2 transition-all duration-200"
                        >
                          <Trash2 className="w-4 h-4" />
                        </Button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
            )}
          </CardContent>
        </Card>

        {/* Footer */}
        {todos.length > 0 && (
          <div className="text-center mt-12 animate-fade-in">
            <div className="inline-flex items-center gap-2 px-6 py-3 bg-white/60 backdrop-blur-sm rounded-full shadow-lg border border-slate-200/50">
              <div className="flex items-center gap-4 text-sm text-slate-600">
                <span className="flex items-center gap-2">
                  <Circle className="w-4 h-4" />
                  Click to complete
                </span>
                <span className="flex items-center gap-2">
                  <Edit3 className="w-4 h-4" />
                  Edit tasks
                </span>
                <span className="flex items-center gap-2">
                  <Trash2 className="w-4 h-4" />
                  Remove tasks
                </span>
              </div>
            </div>
          </div>
        )}
      </div>

      <style>{`
        @keyframes fadeInUp {
          from {
            opacity: 0;
            transform: translateY(20px);
          }
          to {
            opacity: 1;
            transform: translateY(0);
          }
        }
        
        @keyframes fade-in {
          from { opacity: 0; }
          to { opacity: 1; }
        }
        
        @keyframes slide-up {
          from {
            opacity: 0;
            transform: translateY(30px);
          }
          to {
            opacity: 1;
            transform: translateY(0);
          }
        }
        
        .animate-fade-in {
          animation: fade-in 0.6s ease-out;
        }
        
        .animate-slide-up {
          animation: slide-up 0.6s ease-out;
        }
      `}</style>
    </div>
  );
}

export default App;