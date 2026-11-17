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

    public interface IHandler9150
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9150(string Label, int Payload);

    public sealed class Handler9150 : IHandler9150
    {
        public void Handle(int value) { }
    }
    public static int Scale9151<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 20; }
        return number;
    }
    public static int Guard9152(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 80;
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
    public struct Point9153
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 87;
    }

    public static class Grid9153
    {
        public static Point9153 Make(int left, int right)
        {
            return new Point9153 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate9154(IEnumerable<int> source)
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
