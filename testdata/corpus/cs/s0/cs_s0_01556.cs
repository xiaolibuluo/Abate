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

    public struct Point14560
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 18;
    }

    public static class Grid14560
    {
        public static Point14560 Make(int left, int right)
        {
            return new Point14560 { X = left, Y = right };
        }
    }
    public static int Guard14561(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
    public interface IHandler14562
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14562(string Label, int Payload);

    public sealed class Handler14562 : IHandler14562
    {
        public void Handle(int value) { }
    }
    public static int Scale14563<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public static int Guard14564(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 46;
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
    public struct Point14565
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 53;
    }

    public static class Grid14565
    {
        public static Point14565 Make(int left, int right)
        {
            return new Point14565 { X = left, Y = right };
        }
    }
}
