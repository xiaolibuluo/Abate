package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6620")
fun step6620(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 75
}
@JvmName("step6621")
fun step6621(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
class Counter6622 {
    var current: Int by Delegates.observable(89) { _, _, new -> println(new) }

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
object Registry6623 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6623 = "registry-96"
    }
}
@JvmName("step6624")
fun step6624(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
sealed class Outcome6625 {
    data class Success(val payload: String) : Outcome6625()
    data object Failure : Outcome6625()
}

fun describe6625(outcome: Outcome6625): String = when (outcome) {
    is Outcome6625.Success -> outcome.payload
    Outcome6625.Failure -> "failed"
}
