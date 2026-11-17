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

    public static int Scale8730<T>(T value) where T : struct
    {
        int number = 5;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public static int Guard8731(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 12;
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
    public struct Point8732
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 19;
    }

    public static class Grid8732
    {
        public static Point8732 Make(int left, int right)
        {
            return new Point8732 { X = left, Y = right };
        }
    }
    public interface IHandler8733
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8733(string Label, int Payload);

    public sealed class Handler8733 : IHandler8733
    {
        public void Handle(int value) { }
    }
    public struct Point8734
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 33;
    }

    public static class Grid8734
    {
        public static Point8734 Make(int left, int right)
        {
            return new Point8734 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate8735(IEnumerable<int> source)
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
