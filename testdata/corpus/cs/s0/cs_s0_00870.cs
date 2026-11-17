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

    public struct Point8160
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid8160
    {
        public static Point8160 Make(int left, int right)
        {
            return new Point8160 { X = left, Y = right };
        }
    }
    public struct Point8161
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 47;
    }

    public static class Grid8161
    {
        public static Point8161 Make(int left, int right)
        {
            return new Point8161 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate8162(IEnumerable<int> source)
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
    public delegate void Notifier8163(int value);

    public sealed class Source8163
    {
        public event Notifier8163 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler8164
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8164(string Label, int Payload);

    public sealed class Handler8164 : IHandler8164
    {
        public void Handle(int value) { }
    }
}
