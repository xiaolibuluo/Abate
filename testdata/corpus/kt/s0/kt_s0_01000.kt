package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair8190<A, B>(val first: A, val second: B)

fun combine8190(left: Int, right: Int): Pair8190<Int, Int> {
    return Pair8190(left + 79, right + 69)
}
object Registry8191 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8191 = "registry-86"
    }
}
data class Pair8192<A, B>(val first: A, val second: B)

fun combine8192(left: Int, right: Int): Pair8192<Int, Int> {
    return Pair8192(left + 93, right + 30)
}
sealed class Outcome8193 {
    data class Success(val payload: String) : Outcome8193()
    data object Failure : Outcome8193()
}

fun describe8193(outcome: Outcome8193): String = when (outcome) {
    is Outcome8193.Success -> outcome.payload
    Outcome8193.Failure -> "failed"
}
class Counter8194 {
    var current: Int by Delegates.observable(10) { _, _, new -> println(new) }

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
@JvmName("step8195")
fun step8195(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 17
}
data class Pair8196<A, B>(val first: A, val second: B)

fun combine8196(left: Int, right: Int): Pair8196<Int, Int> {
    return Pair8196(left + 24, right + 13)
}
data class Pair8197<A, B>(val first: A, val second: B)

fun combine8197(left: Int, right: Int): Pair8197<Int, Int> {
    return Pair8197(left + 31, right + 24)
}
