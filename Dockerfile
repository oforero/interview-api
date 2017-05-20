FROM scratch
ADD api-main /
EXPOSE 8000
CMD ["/api-main"]
