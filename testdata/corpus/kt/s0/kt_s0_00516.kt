package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step4170")
fun step4170(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
class Counter4171 {
    var current: Int by Delegates.observable(98) { _, _, new -> println(new) }

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
sealed class Outcome4172 {
    data class Success(val payload: String) : Outcome4172()
    data object Failure : Outcome4172()
}

fun describe4172(outcome: Outcome4172): String = when (outcome) {
    is Outcome4172.Success -> outcome.payload
    Outcome4172.Failure -> "failed"
}
object Registry4173 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG4173 = "registry-15"
    }
}
@JvmName("step4174")
fun step4174(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
@JvmName("step4175")
fun step4175(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 29
}
