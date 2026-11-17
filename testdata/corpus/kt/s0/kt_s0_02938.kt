package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24230<A, B>(val first: A, val second: B)

fun combine24230(left: Int, right: Int): Pair24230<Int, Int> {
    return Pair24230(left + 76, right + 12)
}
data class Pair24231<A, B>(val first: A, val second: B)

fun combine24231(left: Int, right: Int): Pair24231<Int, Int> {
    return Pair24231(left + 83, right + 23)
}
object Registry24232 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24232 = "registry-90"
    }
}
fun <T> List<T>.secondOrNull24233(): T? = if (size > 1) this[1] else null

fun render24233(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair24234<A, B>(val first: A, val second: B)

fun combine24234(left: Int, right: Int): Pair24234<Int, Int> {
    return Pair24234(left + 7, right + 56)
}
sealed class Outcome24235 {
    data class Success(val payload: String) : Outcome24235()
    data object Failure : Outcome24235()
}

fun describe24235(outcome: Outcome24235): String = when (outcome) {
    is Outcome24235.Success -> outcome.payload
    Outcome24235.Failure -> "failed"
}
data class Pair24236<A, B>(val first: A, val second: B)

fun combine24236(left: Int, right: Int): Pair24236<Int, Int> {
    return Pair24236(left + 21, right + 17)
}
class Counter24237 {
    var current: Int by Delegates.observable(28) { _, _, new -> println(new) }

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
