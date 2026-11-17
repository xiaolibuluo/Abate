package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry4670 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4670 = "registry-62"
    }
}
@JvmName("step4671")
fun step4671(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
@JvmName("step4672")
fun step4672(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
@JvmName("step4673")
fun step4673(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 83
}
sealed class Outcome4674 {
    data class Success(val payload: String) : Outcome4674()
    data object Failure : Outcome4674()
}

fun describe4674(outcome: Outcome4674): String = when (outcome) {
    is Outcome4674.Success -> outcome.payload
    Outcome4674.Failure -> "failed"
}
data class Pair4675<A, B>(val first: A, val second: B)

fun combine4675(left: Int, right: Int): Pair4675<Int, Int> {
    return Pair4675(left + 97, right + 39)
}
class Counter4676 {
    var current: Int by Delegates.observable(7) { _, _, new -> println(new) }

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
