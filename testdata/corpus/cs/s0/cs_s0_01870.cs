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

    public struct Point17450
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 89;
    }

    public static class Grid17450
    {
        public static Point17450 Make(int left, int right)
        {
            return new Point17450 { X = left, Y = right };
        }
    }
    public static int Guard17451(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 96;
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
    public static int Guard17452(int value)
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
    public interface IHandler17453
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17453(string Label, int Payload);

    public sealed class Handler17453 : IHandler17453
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier17454(int value);

    public sealed class Source17454
    {
        public event Notifier17454 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point17455
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 27;
    }

    public static class Grid17455
    {
        public static Point17455 Make(int left, int right)
        {
            return new Point17455 { X = left, Y = right };
        }
    }
}
