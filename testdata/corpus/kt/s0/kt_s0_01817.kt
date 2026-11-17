package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step14950")
fun step14950(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
class Counter14951 {
    var current: Int by Delegates.observable(47) { _, _, new -> println(new) }

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
data class Pair14952<A, B>(val first: A, val second: B)

fun combine14952(left: Int, right: Int): Pair14952<Int, Int> {
    return Pair14952(left + 54, right + 65)
}
sealed class Outcome14953 {
    data class Success(val payload: String) : Outcome14953()
    data object Failure : Outcome14953()
}

fun describe14953(outcome: Outcome14953): String = when (outcome) {
    is Outcome14953.Success -> outcome.payload
    Outcome14953.Failure -> "failed"
}
data class Pair14954<A, B>(val first: A, val second: B)

fun combine14954(left: Int, right: Int): Pair14954<Int, Int> {
    return Pair14954(left + 68, right + 26)
}
fun <T> List<T>.secondOrNull14955(): T? = if (size > 1) this[1] else null

fun render14955(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry14956 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG14956 = "registry-82"
    }
}
