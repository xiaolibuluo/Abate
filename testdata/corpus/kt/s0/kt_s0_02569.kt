package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair21160<A, B>(val first: A, val second: B)

fun combine21160(left: Int, right: Int): Pair21160<Int, Int> {
    return Pair21160(left + 62, right + 59)
}
class Counter21161 {
    var current: Int by Delegates.observable(69) { _, _, new -> println(new) }

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
sealed class Outcome21162 {
    data class Success(val payload: String) : Outcome21162()
    data object Failure : Outcome21162()
}

fun describe21162(outcome: Outcome21162): String = when (outcome) {
    is Outcome21162.Success -> outcome.payload
    Outcome21162.Failure -> "failed"
}
object Registry21163 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21163 = "registry-83"
    }
}
fun <T, R> List<T>.mapChecked21164(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
