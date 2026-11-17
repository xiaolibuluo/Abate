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

    public static int Guard10590(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 95;
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
    public delegate void Notifier10591(int value);

    public sealed class Source10591
    {
        public event Notifier10591 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler10592
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10592(string Label, int Payload);

    public sealed class Handler10592 : IHandler10592
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier10593(int value);

    public sealed class Source10593
    {
        public event Notifier10593 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point10594
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 26;
    }

    public static class Grid10594
    {
        public static Point10594 Make(int left, int right)
        {
            return new Point10594 { X = left, Y = right };
        }
    }
}
