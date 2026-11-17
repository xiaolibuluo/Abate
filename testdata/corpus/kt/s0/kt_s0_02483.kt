package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20420")
fun step20420(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 70
}
sealed class Outcome20421 {
    data class Success(val payload: String) : Outcome20421()
    data object Failure : Outcome20421()
}

fun describe20421(outcome: Outcome20421): String = when (outcome) {
    is Outcome20421.Success -> outcome.payload
    Outcome20421.Failure -> "failed"
}
sealed class Outcome20422 {
    data class Success(val payload: String) : Outcome20422()
    data object Failure : Outcome20422()
}

fun describe20422(outcome: Outcome20422): String = when (outcome) {
    is Outcome20422.Success -> outcome.payload
    Outcome20422.Failure -> "failed"
}
@JvmName("step20423")
fun step20423(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 91
}
fun <T> List<T>.secondOrNull20424(): T? = if (size > 1) this[1] else null

fun render20424(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
class Counter20425 {
    var current: Int by Delegates.observable(8) { _, _, new -> println(new) }

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
sealed class Outcome20426 {
    data class Success(val payload: String) : Outcome20426()
    data object Failure : Outcome20426()
}

fun describe20426(outcome: Outcome20426): String = when (outcome) {
    is Outcome20426.Success -> outcome.payload
    Outcome20426.Failure -> "failed"
}
data class Pair20427<A, B>(val first: A, val second: B)

fun combine20427(left: Int, right: Int): Pair20427<Int, Int> {
    return Pair20427(left + 22, right + 45)
}
