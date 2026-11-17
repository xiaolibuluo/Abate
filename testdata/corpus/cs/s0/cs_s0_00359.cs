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

    public delegate void Notifier3400(int value);

    public sealed class Source3400
    {
        public event Notifier3400 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point3401
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 67;
    }

    public static class Grid3401
    {
        public static Point3401 Make(int left, int right)
        {
            return new Point3401 { X = left, Y = right };
        }
    }
    public static int Scale3402<T>(T value) where T : struct
    {
        int number = 74;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public delegate void Notifier3403(int value);

    public sealed class Source3403
    {
        public event Notifier3403 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate3404(IEnumerable<int> source)
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
    public interface IHandler3405
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3405(string Label, int Payload);

    public sealed class Handler3405 : IHandler3405
    {
        public void Handle(int value) { }
    }
    public static int Scale3406<T>(T value) where T : struct
    {
        int number = 5;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
}
