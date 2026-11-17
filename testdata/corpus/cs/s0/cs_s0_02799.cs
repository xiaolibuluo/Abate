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

    public async Task<int> Aggregate26180(IEnumerable<int> source)
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
    public struct Point26181
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 96;
    }

    public static class Grid26181
    {
        public static Point26181 Make(int left, int right)
        {
            return new Point26181 { X = left, Y = right };
        }
    }
    public struct Point26182
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 6;
    }

    public static class Grid26182
    {
        public static Point26182 Make(int left, int right)
        {
            return new Point26182 { X = left, Y = right };
        }
    }
    public interface IHandler26183
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26183(string Label, int Payload);

    public sealed class Handler26183 : IHandler26183
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier26184(int value);

    public sealed class Source26184
    {
        public event Notifier26184 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
