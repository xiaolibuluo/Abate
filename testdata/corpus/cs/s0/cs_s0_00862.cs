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

    public struct Point8080
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid8080
    {
        public static Point8080 Make(int left, int right)
        {
            return new Point8080 { X = left, Y = right };
        }
    }
    public interface IHandler8081
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8081(string Label, int Payload);

    public sealed class Handler8081 : IHandler8081
    {
        public void Handle(int value) { }
    }
    public static int Guard8082(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 47;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public async Task<int> Aggregate8083(IEnumerable<int> source)
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
}
