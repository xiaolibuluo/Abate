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

    public static int Guard20780(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 53;
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
    public delegate void Notifier20781(int value);

    public sealed class Source20781
    {
        public event Notifier20781 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier20782(int value);

    public sealed class Source20782
    {
        public event Notifier20782 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler20783
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20783(string Label, int Payload);

    public sealed class Handler20783 : IHandler20783
    {
        public void Handle(int value) { }
    }
    public struct Point20784
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid20784
    {
        public static Point20784 Make(int left, int right)
        {
            return new Point20784 { X = left, Y = right };
        }
    }
}
