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

    public struct Point8720
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid8720
    {
        public static Point8720 Make(int left, int right)
        {
            return new Point8720 { X = left, Y = right };
        }
    }
    public interface IHandler8721
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8721(string Label, int Payload);

    public sealed class Handler8721 : IHandler8721
    {
        public void Handle(int value) { }
    }
    public static int Guard8722(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 6;
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
    public static int Scale8723<T>(T value) where T : struct
    {
        int number = 13;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public static int Guard8724(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 20;
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
    public interface IHandler8725
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8725(string Label, int Payload);

    public sealed class Handler8725 : IHandler8725
    {
        public void Handle(int value) { }
    }
}
