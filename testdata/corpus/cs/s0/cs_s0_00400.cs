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

    public static int Guard3770(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
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
    public static int Scale3771<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public static int Scale3772<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public struct Point3773
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 77;
    }

    public static class Grid3773
    {
        public static Point3773 Make(int left, int right)
        {
            return new Point3773 { X = left, Y = right };
        }
    }
    public interface IHandler3774
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3774(string Label, int Payload);

    public sealed class Handler3774 : IHandler3774
    {
        public void Handle(int value) { }
    }
}
