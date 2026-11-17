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

    public async Task<int> Aggregate3190(IEnumerable<int> source)
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
    public interface IHandler3191
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3191(string Label, int Payload);

    public sealed class Handler3191 : IHandler3191
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier3192(int value);

    public sealed class Source3192
    {
        public event Notifier3192 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point3193
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 99;
    }

    public static class Grid3193
    {
        public static Point3193 Make(int left, int right)
        {
            return new Point3193 { X = left, Y = right };
        }
    }
}
