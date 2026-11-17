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

    public delegate void Notifier6740(int value);

    public sealed class Source6740
    {
        public event Notifier6740 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard6741(int value)
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
    public struct Point6742
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid6742
    {
        public static Point6742 Make(int left, int right)
        {
            return new Point6742 { X = left, Y = right };
        }
    }
    public interface IHandler6743
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6743(string Label, int Payload);

    public sealed class Handler6743 : IHandler6743
    {
        public void Handle(int value) { }
    }
    public interface IHandler6744
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6744(string Label, int Payload);

    public sealed class Handler6744 : IHandler6744
    {
        public void Handle(int value) { }
    }
    public static int Guard6745(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 72;
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
