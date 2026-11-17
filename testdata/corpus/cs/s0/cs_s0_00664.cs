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

    public struct Point6190
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid6190
    {
        public static Point6190 Make(int left, int right)
        {
            return new Point6190 { X = left, Y = right };
        }
    }
    public interface IHandler6191
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6191(string Label, int Payload);

    public sealed class Handler6191 : IHandler6191
    {
        public void Handle(int value) { }
    }
    public static int Scale6192<T>(T value) where T : struct
    {
        int number = 15;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 33; }
        return number;
    }
    public delegate void Notifier6193(int value);

    public sealed class Source6193
    {
        public event Notifier6193 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard6194(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 29;
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
