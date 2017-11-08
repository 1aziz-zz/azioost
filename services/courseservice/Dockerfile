FROM microsoft/dotnet
COPY CourseService.Web /home/CourseService.Web
COPY CourseService.Core /home/CourseService.Core
WORKDIR /home/CourseService.Web
RUN dotnet restore && dotnet publish -c Release -o out
CMD ["dotnet", "out/CourseService.Web.dll"]
