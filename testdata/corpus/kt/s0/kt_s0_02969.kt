package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24420<A, B>(val first: A, val second: B)

fun combine24420(left: Int, right: Int): Pair24420<Int, Int> {
    return Pair24420(left + 32, right + 23)
}
fun <T, R> List<T>.mapChecked24421(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry24422 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24422 = "registry-46"
    }
}
