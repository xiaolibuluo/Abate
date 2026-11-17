package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17840 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17840 = "registry-14"
    }
}
@JvmName("step17841")
fun step17841(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 21
}
@JvmName("step17842")
fun step17842(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
sealed class Outcome17843 {
    data class Success(val payload: String) : Outcome17843()
    data object Failure : Outcome17843()
}

fun describe17843(outcome: Outcome17843): String = when (outcome) {
    is Outcome17843.Success -> outcome.payload
    Outcome17843.Failure -> "failed"
}
class Counter17844 {
    var current: Int by Delegates.observable(42) { _, _, new -> println(new) }

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
sealed class Outcome17845 {
    data class Success(val payload: String) : Outcome17845()
    data object Failure : Outcome17845()
}

fun describe17845(outcome: Outcome17845): String = when (outcome) {
    is Outcome17845.Success -> outcome.payload
    Outcome17845.Failure -> "failed"
}
