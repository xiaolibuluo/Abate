package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25540 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25540 = "registry-33"
    }
}
data class Pair25541<A, B>(val first: A, val second: B)

fun combine25541(left: Int, right: Int): Pair25541<Int, Int> {
    return Pair25541(left + 40, right + 25)
}
fun <T> List<T>.secondOrNull25542(): T? = if (size > 1) this[1] else null

fun render25542(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter25543 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

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
sealed class Outcome25544 {
    data class Success(val payload: String) : Outcome25544()
    data object Failure : Outcome25544()
}

fun describe25544(outcome: Outcome25544): String = when (outcome) {
    is Outcome25544.Success -> outcome.payload
    Outcome25544.Failure -> "failed"
}
