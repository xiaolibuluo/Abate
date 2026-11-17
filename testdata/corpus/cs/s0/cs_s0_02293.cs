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

    public struct Point21400
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid21400
    {
        public static Point21400 Make(int left, int right)
        {
            return new Point21400 { X = left, Y = right };
        }
    }
    public delegate void Notifier21401(int value);

    public sealed class Source21401
    {
        public event Notifier21401 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate21402(IEnumerable<int> source)
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
    public delegate void Notifier21403(int value);

    public sealed class Source21403
    {
        public event Notifier21403 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21404
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21404(string Label, int Payload);

    public sealed class Handler21404 : IHandler21404
    {
        public void Handle(int value) { }
    }
    public struct Point21405
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid21405
    {
        public static Point21405 Make(int left, int right)
        {
            return new Point21405 { X = left, Y = right };
        }
    }
}
