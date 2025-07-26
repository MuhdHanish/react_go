# ReactGo Todo - Full Stack Todo Application

A modern, full-stack todo application featuring a sleek React frontend with a robust Go Fiber + MongoDB backend API.

## 🌟 Overview

ReactGo Todo combines a beautiful, animated React frontend with a powerful Go Fiber backend to deliver a premium task management experience. The application features real-time updates, smooth animations, and professional design aesthetics.

## 🚀 Features

### Frontend (React + TypeScript)
- **Modern UI/UX** - Sleek, professional design with glass morphism effects
- **Rich Animations** - Smooth transitions, hover effects, and micro-interactions
- **Real-time Stats** - Live dashboard showing task completion metrics
- **Responsive Design** - Optimized for desktop, tablet, and mobile devices
- **Toast Notifications** - User feedback for all actions
- **Inline Editing** - Edit tasks directly in the interface
- **Progress Tracking** - Visual completion percentage and statistics

### Backend (Go Fiber + MongoDB)
- **Full CRUD Operations** - Create, Read, Update, Delete todos
- **MongoDB Integration** - Persistent data storage with MongoDB Atlas
- **RESTful API Design** - Clean and consistent API endpoints
- **Error Handling** - Comprehensive error responses
- **Environment Configuration** - Secure configuration management
- **Modular Architecture** - Well-organized code structure

## 🛠️ Tech Stack

### Frontend
- **Framework**: React 18 with TypeScript
- **Styling**: Tailwind CSS with custom animations
- **UI Components**: shadcn/ui component library
- **Notifications**: Sonner toast notifications
- **Icons**: Lucide React icons
- **Build Tool**: Vite
- **HTTP Client**: Fetch API

### Backend
- **Framework**: [Fiber](https://gofiber.io/) (Go)
- **Database**: MongoDB Atlas
- **Language**: Go 1.21+
- **ODM**: MongoDB Go Driver

## 📋 Prerequisites

- Node.js 18+ and npm/yarn/pnpm
- Go 1.21 or higher
- MongoDB Atlas account (or local MongoDB)
- Git

## ⚡ Quick Start

### 1. Clone the repository
```bash
git clone https://github.com/MuhdHanish/react_go.git
cd react_go
```

### 2. Backend Setup

#### Install Go dependencies
```bash
go mod tidy
```

#### Environment Setup
Create a `.env` file in the root directory:
```env
PORT=8000
MONGODB_URI="mongodb+srv://username:password@cluster0.xxxxx.mongodb.net/react_go?retryWrites=true&w=majority&appName=Cluster0"
```

#### Start the backend server
```bash
go run main.go
```

The API server will start on `http://localhost:8000`

### 3. Frontend Setup

#### Navigate to client directory
```bash
cd client  # or wherever your React app is located
```

#### Install dependencies
```bash
npm install
# or
yarn install
# or
pnpm install
```

#### Start the development server
```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

The React app will start on `http://localhost:5173` (or similar)

## 📁 Project Structure

```
react_go/
├── main.go                 # Go application entry point
├── .env                   # Environment variables
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies
├── config/
│   └── database.go        # Database configuration
├── models/
│   └── todo.go           # Todo data models
├── handlers/
│   └── todo_handler.go   # HTTP request handlers
├── routes/
│   └── todo_routes.go    # Route definitions
├── utils/
│   └── response.go       # Response utilities
└── client/                # React frontend
    ├── public/
    ├── src/
    │   ├── components/    # React components
    │   ├── lib/          # Utility functions
    │   ├── types/        # TypeScript type definitions
    │   ├── App.tsx       # Main React component
    │   └── main.tsx      # React entry point
    ├── package.json      # Frontend dependencies
    ├── tailwind.config.js # Tailwind configuration
    ├── tsconfig.json     # TypeScript configuration
    └── vite.config.ts    # Vite configuration
```

## 🎨 Frontend Features

### UI Components
- **ReactGo Header** - Animated logo with gradient effects
- **Stats Dashboard** - Real-time metrics with hover animations
- **Task Input** - Elegant input field with focus effects
- **Task List** - Animated task items with smooth interactions
- **Action Buttons** - Hover effects and loading states

### Animations & Interactions
- **Page Load Animations** - Staggered component reveals
- **Hover Effects** - Scale, shadow, and color transitions
- **Task Completion** - Smooth check animations
- **Loading States** - Custom spinners and skeleton screens
- **Toast Notifications** - Success/error feedback

### Responsive Design
- **Mobile Optimized** - Touch-friendly interfaces
- **Tablet Support** - Adaptive layouts
- **Desktop Experience** - Full feature accessibility

## 🔗 API Endpoints

### Base URL: `http://localhost:8000/api`

| Method | Endpoint | Description | Frontend Usage |
|--------|----------|-------------|----------------|
| `GET` | `/todos` | Get all todos | Load task list |
| `GET` | `/todos/:id` | Get a specific todo | Individual task data |
| `POST` | `/todos` | Create a new todo | Add new task |
| `PUT` | `/todos/:id` | Update a todo | Edit task content |
| `PATCH` | `/todos/:id` | Toggle todo completion | Mark complete/incomplete |
| `DELETE` | `/todos/:id` | Delete a todo | Remove task |

## 📝 Frontend API Integration

The React frontend communicates with the backend through these key functions:

### Task Management
```typescript
// Fetch all todos
const fetchTodos = async () => {
  const response = await fetch(`${API_BASE}/todos`);
  const data = await response.json();
  return data;
};

// Create new todo
const createTodo = async (body: string) => {
  const response = await fetch(`${API_BASE}/todos`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ body })
  });
  return response.json();
};

// Complete todo
const completeTodo = async (id: string) => {
  const response = await fetch(`${API_BASE}/todos/${id}`, {
    method: 'PATCH'
  });
  return response.json();
};
```

## 🎯 User Experience Features

### Real-time Feedback
- **Instant Updates** - UI updates immediately on actions
- **Loading States** - Visual feedback during API calls
- **Error Handling** - User-friendly error messages
- **Success Notifications** - Confirmation of completed actions

### Accessibility
- **Keyboard Navigation** - Full keyboard support
- **Screen Reader Support** - Semantic HTML structure
- **Color Contrast** - WCAG compliant color schemes
- **Focus Management** - Clear focus indicators

## 📊 Response Format

All API responses follow this consistent format:

### Success Response
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

### Error Response
```json
{
  "success": false,
  "message": "Error description",
  "error": "Detailed error message"
}
```

## 🗃️ Data Model

### Todo Schema
```go
type Todo struct {
    ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Completed bool               `json:"completed" bson:"completed"`
    Body      string             `json:"body" bson:"body"`
}
```

### TypeScript Interface
```typescript
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
```

## 🔧 Development

### Backend Development
```bash
# Install Air for live reloading
go install github.com/air-verse/air@latest

# Start development server with hot reload
air
```

### Frontend Development
```bash
# Start development server
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

### Environment Variables

#### Backend (.env)
| Variable | Description | Required |
|----------|-------------|----------|
| `PORT` | Server port (default: 8000) | No |
| `MONGODB_URI` | MongoDB connection string | Yes |

#### Frontend
The frontend is configured to connect to `http://localhost:8000/api` by default. Update the `API_BASE` constant in your React app to match your backend URL.

## 🚦 HTTP Status Codes

- `200` - OK (Success)
- `201` - Created (Resource created successfully)
- `400` - Bad Request (Invalid input)
- `404` - Not Found (Resource not found)  
- `500` - Internal Server Error (Server error)

## 🧪 Testing

### Backend Testing
Test the API endpoints using:
- [Postman](https://www.postman.com/)
- [Insomnia](https://insomnia.rest/)
- [Thunder Client](https://www.thunderclient.com/) (VS Code extension)
- cURL (command line)

### Frontend Testing
```bash
# Run tests (if configured)
npm run test

# Type checking
npm run type-check

# Linting
npm run lint
```

## 📱 Deployment

### Frontend Deployment
```bash
# Build for production
npm run build

# The dist/ folder contains the built application
# Deploy to Vercel, Netlify, or any static hosting service
```

### Backend Deployment
- Deploy to platforms like Railway, Heroku, or DigitalOcean
- Ensure environment variables are configured
- Update CORS settings for production frontend URL

## 🔒 Error Handling

### Backend Error Handling
- Invalid JSON payload
- Missing required fields
- Invalid MongoDB ObjectID format
- Database connection errors
- Resource not found scenarios

### Frontend Error Handling
- Network connectivity issues
- API response errors
- Form validation errors
- Loading state management
- User-friendly error messages

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Muhammed Hanish** - [@MuhdHanish](https://github.com/MuhdHanish)

## 🙏 Acknowledgments

- [Go Fiber](https://gofiber.io/) - Web framework
- [React](https://reactjs.org/) - Frontend library
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS
- [shadcn/ui](https://ui.shadcn.com/) - Component library
- [Lucide](https://lucide.dev/) - Icon library

⭐ If you found this project helpful, please give it a star!
