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

    public static int Scale10170<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public static int Guard10171(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 41;
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
    public interface IHandler10172
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10172(string Label, int Payload);

    public sealed class Handler10172 : IHandler10172
    {
        public void Handle(int value) { }
    }
    public struct Point10173
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 55;
    }

    public static class Grid10173
    {
        public static Point10173 Make(int left, int right)
        {
            return new Point10173 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate10174(IEnumerable<int> source)
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
