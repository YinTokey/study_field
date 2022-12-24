package com.example.helloworld.servlet;

import jakarta.servlet.annotation.WebServlet;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import jakarta.servlet.http.HttpServlet;

@WebServlet(name="helloServlet", urlPatterns = "/helloServlet")
public class HelloServlet extends HttpServlet {
    // @Override
    protected void doGet(HttpServletRequest request, HttpServletResponse response) {
        System.out.println("Running Hello Servlet doGet method");
    }
}
