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

    public struct Point7280
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid7280
    {
        public static Point7280 Make(int left, int right)
        {
            return new Point7280 { X = left, Y = right };
        }
    }
    public delegate void Notifier7281(int value);

    public sealed class Source7281
    {
        public event Notifier7281 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard7282(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 74;
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
    public struct Point7283
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid7283
    {
        public static Point7283 Make(int left, int right)
        {
            return new Point7283 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate7284(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public delegate void Notifier7285(int value);

    public sealed class Source7285
    {
        public event Notifier7285 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
