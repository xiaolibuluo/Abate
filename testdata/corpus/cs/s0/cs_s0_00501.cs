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

    public interface IHandler4710
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4710(string Label, int Payload);

    public sealed class Handler4710 : IHandler4710
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate4711(IEnumerable<int> source)
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
    public delegate void Notifier4712(int value);

    public sealed class Source4712
    {
        public event Notifier4712 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier4713(int value);

    public sealed class Source4713
    {
        public event Notifier4713 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point4714
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 45;
    }

    public static class Grid4714
    {
        public static Point4714 Make(int left, int right)
        {
            return new Point4714 { X = left, Y = right };
        }
    }
}
