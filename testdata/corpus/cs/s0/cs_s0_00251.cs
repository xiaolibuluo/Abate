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

    public static int Scale2380<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
    public interface IHandler2381
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2381(string Label, int Payload);

    public sealed class Handler2381 : IHandler2381
    {
        public void Handle(int value) { }
    }
    public struct Point2382
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid2382
    {
        public static Point2382 Make(int left, int right)
        {
            return new Point2382 { X = left, Y = right };
        }
    }
    public static int Guard2383(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 16;
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
    public delegate void Notifier2384(int value);

    public sealed class Source2384
    {
        public event Notifier2384 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point2385
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 30;
    }

    public static class Grid2385
    {
        public static Point2385 Make(int left, int right)
        {
            return new Point2385 { X = left, Y = right };
        }
    }
}
