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

    public struct Point2300
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid2300
    {
        public static Point2300 Make(int left, int right)
        {
            return new Point2300 { X = left, Y = right };
        }
    }
    public static int Scale2301<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public struct Point2302
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid2302
    {
        public static Point2302 Make(int left, int right)
        {
            return new Point2302 { X = left, Y = right };
        }
    }
    public static int Guard2303(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 9;
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
    public interface IHandler2304
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope2304(string Label, int Payload);

    public sealed class Handler2304 : IHandler2304
    {
        public void Handle(int value) { }
    }
}
