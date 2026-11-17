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

    public static int Guard3790(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 82;
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
    public struct Point3791
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid3791
    {
        public static Point3791 Make(int left, int right)
        {
            return new Point3791 { X = left, Y = right };
        }
    }
    public interface IHandler3792
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3792(string Label, int Payload);

    public sealed class Handler3792 : IHandler3792
    {
        public void Handle(int value) { }
    }
    public static int Scale3793<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public struct Point3794
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 13;
    }

    public static class Grid3794
    {
        public static Point3794 Make(int left, int right)
        {
            return new Point3794 { X = left, Y = right };
        }
    }
}
