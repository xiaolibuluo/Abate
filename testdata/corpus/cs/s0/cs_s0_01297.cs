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

    public delegate void Notifier12150(int value);

    public sealed class Source12150
    {
        public event Notifier12150 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point12151
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 93;
    }

    public static class Grid12151
    {
        public static Point12151 Make(int left, int right)
        {
            return new Point12151 { X = left, Y = right };
        }
    }
    public struct Point12152
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 100;
    }

    public static class Grid12152
    {
        public static Point12152 Make(int left, int right)
        {
            return new Point12152 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate12153(IEnumerable<int> source)
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
    public interface IHandler12154
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12154(string Label, int Payload);

    public sealed class Handler12154 : IHandler12154
    {
        public void Handle(int value) { }
    }
    public static int Scale12155<T>(T value) where T : struct
    {
        int number = 24;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 29; }
        return number;
    }
}
