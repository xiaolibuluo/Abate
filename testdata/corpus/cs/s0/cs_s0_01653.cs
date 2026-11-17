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

    public interface IHandler15470
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15470(string Label, int Payload);

    public sealed class Handler15470 : IHandler15470
    {
        public void Handle(int value) { }
    }
    public static int Guard15471(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 44;
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
    public interface IHandler15472
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15472(string Label, int Payload);

    public sealed class Handler15472 : IHandler15472
    {
        public void Handle(int value) { }
    }
    public static int Scale15473<T>(T value) where T : struct
    {
        int number = 58;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public static int Scale15474<T>(T value) where T : struct
    {
        int number = 65;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public struct Point15475
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 72;
    }

    public static class Grid15475
    {
        public static Point15475 Make(int left, int right)
        {
            return new Point15475 { X = left, Y = right };
        }
    }
}
