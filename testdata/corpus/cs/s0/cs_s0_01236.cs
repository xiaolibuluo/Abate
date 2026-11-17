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

    public async Task<int> Aggregate11590(IEnumerable<int> source)
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
    public struct Point11591
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 44;
    }

    public static class Grid11591
    {
        public static Point11591 Make(int left, int right)
        {
            return new Point11591 { X = left, Y = right };
        }
    }
    public interface IHandler11592
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11592(string Label, int Payload);

    public sealed class Handler11592 : IHandler11592
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier11593(int value);

    public sealed class Source11593
    {
        public event Notifier11593 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale11594<T>(T value) where T : struct
    {
        int number = 65;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
}
