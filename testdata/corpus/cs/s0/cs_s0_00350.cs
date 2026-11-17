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

    public struct Point3320
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid3320
    {
        public static Point3320 Make(int left, int right)
        {
            return new Point3320 { X = left, Y = right };
        }
    }
    public struct Point3321
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid3321
    {
        public static Point3321 Make(int left, int right)
        {
            return new Point3321 { X = left, Y = right };
        }
    }
    public static int Guard3322(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 67;
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
    public interface IHandler3323
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3323(string Label, int Payload);

    public sealed class Handler3323 : IHandler3323
    {
        public void Handle(int value) { }
    }
    public static int Scale3324<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
}
