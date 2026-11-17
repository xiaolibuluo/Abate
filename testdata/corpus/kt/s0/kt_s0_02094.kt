package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull17290(): T? = if (size > 1) this[1] else null

fun render17290(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17291 {
    data class Success(val payload: String) : Outcome17291()
    data object Failure : Outcome17291()
}

fun describe17291(outcome: Outcome17291): String = when (outcome) {
    is Outcome17291.Success -> outcome.payload
    Outcome17291.Failure -> "failed"
}
data class Pair17292<A, B>(val first: A, val second: B)

fun combine17292(left: Int, right: Int): Pair17292<Int, Int> {
    return Pair17292(left + 89, right + 56)
}
object Registry17293 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17293 = "registry-96"
    }
}
class Counter17294 {
    var current: Int by Delegates.observable(6) { _, _, new -> println(new) }

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
