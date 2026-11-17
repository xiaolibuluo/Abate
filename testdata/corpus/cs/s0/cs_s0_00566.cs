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

    public static int Scale5310<T>(T value) where T : struct
    {
        int number = 21;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 66; }
        return number;
    }
    public interface IHandler5311
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5311(string Label, int Payload);

    public sealed class Handler5311 : IHandler5311
    {
        public void Handle(int value) { }
    }
    public static int Scale5312<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public delegate void Notifier5313(int value);

    public sealed class Source5313
    {
        public event Notifier5313 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier5314(int value);

    public sealed class Source5314
    {
        public event Notifier5314 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point5315
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 56;
    }

    public static class Grid5315
    {
        public static Point5315 Make(int left, int right)
        {
            return new Point5315 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate5316(IEnumerable<int> source)
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
