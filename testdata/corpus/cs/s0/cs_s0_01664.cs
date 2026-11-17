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

    public async Task<int> Aggregate15560(IEnumerable<int> source)
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
    public static int Scale15561<T>(T value) where T : struct
    {
        int number = 64;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 54; }
        return number;
    }
    public interface IHandler15562
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15562(string Label, int Payload);

    public sealed class Handler15562 : IHandler15562
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier15563(int value);

    public sealed class Source15563
    {
        public event Notifier15563 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point15564
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 85;
    }

    public static class Grid15564
    {
        public static Point15564 Make(int left, int right)
        {
            return new Point15564 { X = left, Y = right };
        }
    }
    public delegate void Notifier15565(int value);

    public sealed class Source15565
    {
        public event Notifier15565 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate15566(IEnumerable<int> source)
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
}
