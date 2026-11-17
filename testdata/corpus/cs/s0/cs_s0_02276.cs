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

    public async Task<int> Aggregate21260(IEnumerable<int> source)
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
    public struct Point21261
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid21261
    {
        public static Point21261 Make(int left, int right)
        {
            return new Point21261 { X = left, Y = right };
        }
    }
    public static int Scale21262<T>(T value) where T : struct
    {
        int number = 12;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
    public delegate void Notifier21263(int value);

    public sealed class Source21263
    {
        public event Notifier21263 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21264
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21264(string Label, int Payload);

    public sealed class Handler21264 : IHandler21264
    {
        public void Handle(int value) { }
    }
}
