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

    public struct Point22260
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 37;
    }

    public static class Grid22260
    {
        public static Point22260 Make(int left, int right)
        {
            return new Point22260 { X = left, Y = right };
        }
    }
    public interface IHandler22261
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22261(string Label, int Payload);

    public sealed class Handler22261 : IHandler22261
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22262(int value);

    public sealed class Source22262
    {
        public event Notifier22262 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22263<T>(T value) where T : struct
    {
        int number = 58;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
    public static int Guard22264(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 65;
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
}
