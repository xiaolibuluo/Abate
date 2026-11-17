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

    public delegate void Notifier11570(int value);

    public sealed class Source11570
    {
        public event Notifier11570 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate11571(IEnumerable<int> source)
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
    public struct Point11572
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid11572
    {
        public static Point11572 Make(int left, int right)
        {
            return new Point11572 { X = left, Y = right };
        }
    }
    public interface IHandler11573
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11573(string Label, int Payload);

    public sealed class Handler11573 : IHandler11573
    {
        public void Handle(int value) { }
    }
    public struct Point11574
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid11574
    {
        public static Point11574 Make(int left, int right)
        {
            return new Point11574 { X = left, Y = right };
        }
    }
}
