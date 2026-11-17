package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair20720<A, B>(val first: A, val second: B)

fun combine20720(left: Int, right: Int): Pair20720<Int, Int> {
    return Pair20720(left + 72, right + 56)
}
sealed class Outcome20721 {
    data class Success(val payload: String) : Outcome20721()
    data object Failure : Outcome20721()
}

fun describe20721(outcome: Outcome20721): String = when (outcome) {
    is Outcome20721.Success -> outcome.payload
    Outcome20721.Failure -> "failed"
}
object Registry20722 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20722 = "registry-86"
    }
}
@JvmName("step20723")
fun step20723(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
class Counter20724 {
    var current: Int by Delegates.observable(100) { _, _, new -> println(new) }

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
@JvmName("step20725")
fun step20725(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 10
}
