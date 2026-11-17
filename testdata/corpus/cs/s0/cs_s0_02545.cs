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

    public interface IHandler23780
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23780(string Label, int Payload);

    public sealed class Handler23780 : IHandler23780
    {
        public void Handle(int value) { }
    }
    public struct Point23781
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid23781
    {
        public static Point23781 Make(int left, int right)
        {
            return new Point23781 { X = left, Y = right };
        }
    }
    public static int Guard23782(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 87;
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
    public interface IHandler23783
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23783(string Label, int Payload);

    public sealed class Handler23783 : IHandler23783
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier23784(int value);

    public sealed class Source23784
    {
        public event Notifier23784 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler23785
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23785(string Label, int Payload);

    public sealed class Handler23785 : IHandler23785
    {
        public void Handle(int value) { }
    }
}
