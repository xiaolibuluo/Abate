using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public interface IHandler15440
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15440(string Label, int Payload);

    public sealed class Handler15440 : IHandler15440
    {
        public void Handle(int value) { }
    }
    public struct Point15441
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid15441
    {
        public static Point15441 Make(int left, int right)
        {
            return new Point15441 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate15442(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public struct Point15443
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 19;
    }

    public static class Grid15443
    {
        public static Point15443 Make(int left, int right)
        {
            return new Point15443 { X = left, Y = right };
        }
    }
}
