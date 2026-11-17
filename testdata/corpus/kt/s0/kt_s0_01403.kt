package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair11550<A, B>(val first: A, val second: B)

fun combine11550(left: Int, right: Int): Pair11550<Int, Int> {
    return Pair11550(left + 82, right + 42)
}
object Registry11551 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11551 = "registry-89"
    }
}
data class Pair11552<A, B>(val first: A, val second: B)

fun combine11552(left: Int, right: Int): Pair11552<Int, Int> {
    return Pair11552(left + 96, right + 64)
}
sealed class Outcome11553 {
    data class Success(val payload: String) : Outcome11553()
    data object Failure : Outcome11553()
}

fun describe11553(outcome: Outcome11553): String = when (outcome) {
    is Outcome11553.Success -> outcome.payload
    Outcome11553.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked11554(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11555<A, B>(val first: A, val second: B)

fun combine11555(left: Int, right: Int): Pair11555<Int, Int> {
    return Pair11555(left + 20, right + 36)
}
sealed class Outcome11556 {
    data class Success(val payload: String) : Outcome11556()
    data object Failure : Outcome11556()
}

fun describe11556(outcome: Outcome11556): String = when (outcome) {
    is Outcome11556.Success -> outcome.payload
    Outcome11556.Failure -> "failed"
}
