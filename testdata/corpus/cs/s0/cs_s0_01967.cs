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

    public delegate void Notifier18330(int value);

    public sealed class Source18330
    {
        public event Notifier18330 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18331<T>(T value) where T : struct
    {
        int number = 76;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public struct Point18332
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid18332
    {
        public static Point18332 Make(int left, int right)
        {
            return new Point18332 { X = left, Y = right };
        }
    }
    public static int Guard18333(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 90;
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
    public interface IHandler18334
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18334(string Label, int Payload);

    public sealed class Handler18334 : IHandler18334
    {
        public void Handle(int value) { }
    }
}
