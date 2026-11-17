package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter20220 {
    var current: Int by Delegates.observable(101) { _, _, new -> println(new) }

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
fun <T, R> List<T>.mapChecked20221(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull20222(): T? = if (size > 1) this[1] else null

fun render20222(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20223<A, B>(val first: A, val second: B)

fun combine20223(left: Int, right: Int): Pair20223<Int, Int> {
    return Pair20223(left + 25, right + 44)
}
sealed class Outcome20224 {
    data class Success(val payload: String) : Outcome20224()
    data object Failure : Outcome20224()
}

fun describe20224(outcome: Outcome20224): String = when (outcome) {
    is Outcome20224.Success -> outcome.payload
    Outcome20224.Failure -> "failed"
}
