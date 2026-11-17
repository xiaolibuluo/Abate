package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step20740")
fun step20740(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
fun <T, R> List<T>.mapChecked20741(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry20742 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20742 = "registry-15"
    }
}
sealed class Outcome20743 {
    data class Success(val payload: String) : Outcome20743()
    data object Failure : Outcome20743()
}

fun describe20743(outcome: Outcome20743): String = when (outcome) {
    is Outcome20743.Success -> outcome.payload
    Outcome20743.Failure -> "failed"
}
sealed class Outcome20744 {
    data class Success(val payload: String) : Outcome20744()
    data object Failure : Outcome20744()
}

fun describe20744(outcome: Outcome20744): String = when (outcome) {
    is Outcome20744.Success -> outcome.payload
    Outcome20744.Failure -> "failed"
}
sealed class Outcome20745 {
    data class Success(val payload: String) : Outcome20745()
    data object Failure : Outcome20745()
}

fun describe20745(outcome: Outcome20745): String = when (outcome) {
    is Outcome20745.Success -> outcome.payload
    Outcome20745.Failure -> "failed"
}
