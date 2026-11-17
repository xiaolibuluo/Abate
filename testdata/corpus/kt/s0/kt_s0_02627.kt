package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter21680 {
    var current: Int by Delegates.observable(59) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
class Counter21681 {
    var current: Int by Delegates.observable(66) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
object Registry21682 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21682 = "registry-73"
    }
}
data class Pair21683<A, B>(val first: A, val second: B)

fun combine21683(left: Int, right: Int): Pair21683<Int, Int> {
    return Pair21683(left + 80, right + 29)
}
sealed class Outcome21684 {
    data class Success(val payload: String) : Outcome21684()
    data object Failure : Outcome21684()
}

fun describe21684(outcome: Outcome21684): String = when (outcome) {
    is Outcome21684.Success -> outcome.payload
    Outcome21684.Failure -> "failed"
}
