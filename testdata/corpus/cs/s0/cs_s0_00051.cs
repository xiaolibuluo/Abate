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

    public delegate void Notifier480(int value);

    public sealed class Source480
    {
        public event Notifier480 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler481
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope481(string Label, int Payload);

    public sealed class Handler481 : IHandler481
    {
        public void Handle(int value) { }
    }
    public struct Point482
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 61;
    }

    public static class Grid482
    {
        public static Point482 Make(int left, int right)
        {
            return new Point482 { X = left, Y = right };
        }
    }
    public delegate void Notifier483(int value);

    public sealed class Source483
    {
        public event Notifier483 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale484<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public async Task<int> Aggregate485(IEnumerable<int> source)
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
    public static int Guard486(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 89;
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
