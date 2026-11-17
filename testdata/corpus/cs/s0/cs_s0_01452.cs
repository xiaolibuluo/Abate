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

    public interface IHandler13610
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13610(string Label, int Payload);

    public sealed class Handler13610 : IHandler13610
    {
        public void Handle(int value) { }
    }
    public struct Point13611
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid13611
    {
        public static Point13611 Make(int left, int right)
        {
            return new Point13611 { X = left, Y = right };
        }
    }
    public interface IHandler13612
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13612(string Label, int Payload);

    public sealed class Handler13612 : IHandler13612
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier13613(int value);

    public sealed class Source13613
    {
        public event Notifier13613 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate13614(IEnumerable<int> source)
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
    public static int Scale13615<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
}
