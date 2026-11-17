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

    public struct Point26250
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 83;
    }

    public static class Grid26250
    {
        public static Point26250 Make(int left, int right)
        {
            return new Point26250 { X = left, Y = right };
        }
    }
    public delegate void Notifier26251(int value);

    public sealed class Source26251
    {
        public event Notifier26251 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate26252(IEnumerable<int> source)
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
    public interface IHandler26253
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26253(string Label, int Payload);

    public sealed class Handler26253 : IHandler26253
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier26254(int value);

    public sealed class Source26254
    {
        public event Notifier26254 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point26255
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid26255
    {
        public static Point26255 Make(int left, int right)
        {
            return new Point26255 { X = left, Y = right };
        }
    }
}
