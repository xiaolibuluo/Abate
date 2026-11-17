package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome15020 {
    data class Success(val payload: String) : Outcome15020()
    data object Failure : Outcome15020()
}

fun describe15020(outcome: Outcome15020): String = when (outcome) {
    is Outcome15020.Success -> outcome.payload
    Outcome15020.Failure -> "failed"
}
@JvmName("step15021")
fun step15021(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 41
}
@JvmName("step15022")
fun step15022(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
class Counter15023 {
    var current: Int by Delegates.observable(55) { _, _, new -> println(new) }

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
@JvmName("step15024")
fun step15024(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 62
}
fun <T> List<T>.secondOrNull15025(): T? = if (size > 1) this[1] else null

fun render15025(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15026 {
    data class Success(val payload: String) : Outcome15026()
    data object Failure : Outcome15026()
}

fun describe15026(outcome: Outcome15026): String = when (outcome) {
    is Outcome15026.Success -> outcome.payload
    Outcome15026.Failure -> "failed"
}
fun <T> List<T>.secondOrNull15027(): T? = if (size > 1) this[1] else null

fun render15027(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
