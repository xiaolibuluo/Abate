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

    public static int Guard21880(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 28;
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
    public interface IHandler21881
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21881(string Label, int Payload);

    public sealed class Handler21881 : IHandler21881
    {
        public void Handle(int value) { }
    }
    public struct Point21882
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 42;
    }

    public static class Grid21882
    {
        public static Point21882 Make(int left, int right)
        {
            return new Point21882 { X = left, Y = right };
        }
    }
    public struct Point21883
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 49;
    }

    public static class Grid21883
    {
        public static Point21883 Make(int left, int right)
        {
            return new Point21883 { X = left, Y = right };
        }
    }
    public static int Scale21884<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 58; }
        return number;
    }
    public static int Scale21885<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
}
